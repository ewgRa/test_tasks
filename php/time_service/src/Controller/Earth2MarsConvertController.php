<?php

namespace App\Controller;

use App\Mars\Earth2MsdConverter;
use App\Mars\MarsDateTime;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\Form\Extension\Core\Type\DateTimeType;
use Symfony\Component\Form\FormError;
use Symfony\Component\Form\FormInterface;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Validator\Constraints\NotBlank;

class Earth2MarsConvertController extends AbstractController
{
    private Earth2MsdConverter $msdConverter;

    public function __construct(Earth2MsdConverter $msdConverter)
    {
        $this->msdConverter = $msdConverter;
    }

    /**
     * @Route("/v1/earth2mars")
     */
    public function convert(Request $request): JsonResponse
    {
        $form = $this->createRequestForm();

        if (!$this->submit($request, $form)) {
            return $this->badRequestResponse($form);
        }

        /** @var \DateTime $datetime */
        $datetime = $form->get('datetime')->getData();
        $marsDatetime = MarsDateTime::createFromMsd($this->msdConverter->convert($datetime));

        return new JsonResponse([
            'msd' => (float)$marsDatetime->format('S'),
            'mtc' => $marsDatetime->format('H:i:s')
        ]);
    }

    private function createRequestForm(): FormInterface
    {
        return $this->createFormBuilder(null, ['csrf_protection' => false])
            ->add('datetime', DateTimeType::class, [
                'html5' => false,
                'widget' => 'single_text',
                'format' => "yyyy-MM-dd'T'HH:mm:ssZ",
                'constraints' => [
                    new NotBlank(),
                ],
            ])
            ->getForm()
            ;
    }

    private function submit(Request $request, FormInterface $form): bool
    {
        try {
            $jsonRequest = json_decode($request->getContent(), true, 512, JSON_THROW_ON_ERROR);
            $form->submit($jsonRequest);
        } catch (\JsonException $e) {
            $form->addError(new FormError("Request cannot be parsed as JSON"));
            return false;
        }

        return $form->isValid();
    }

    private function badRequestResponse(FormInterface $form): JsonResponse
    {
        $errors = [];

        foreach ($form->getErrors(true, true) as $error) {
            $errors[$error->getOrigin()->getName()][] = $error->getMessage();
        }

        return new JsonResponse($errors, Response::HTTP_BAD_REQUEST);
    }
}
