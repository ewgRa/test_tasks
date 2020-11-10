<?php

namespace App\Tests\Controller;

use Symfony\Bundle\FrameworkBundle\KernelBrowser;
use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

class Earth2MarsConvertControllerTest extends WebTestCase
{
    public function testSuccess()
    {
        $client = static::createClient();
        $this->sendRequest($client, '{"datetime": "2120-01-04T01:00:00+0100"}');
        $this->assertEquals(Response::HTTP_OK, $client->getResponse()->getStatusCode());
        $this->assertEquals('{"msd":87450.377581478,"mtc":"09:03:43"}', $client->getResponse()->getContent());
    }

    public function testAuthRequired()
    {
        $client = static::createClient();
        $client->request(Request::METHOD_POST, '/v1/earth2mars');
        $this->assertEquals(Response::HTTP_UNAUTHORIZED, $client->getResponse()->getStatusCode());
    }

    /**
     * @dataProvider badRequestCases
     */
    public function testBadRequest(string $content, string $expectedResponse, string $case)
    {
        $client = static::createClient();
        $this->sendRequest($client, $content);

        $this->assertEquals(Response::HTTP_BAD_REQUEST, $client->getResponse()->getStatusCode());
        $this->assertEquals($expectedResponse, $client->getResponse()->getContent(), $case);
    }

    public function badRequestCases(): array
    {
        return [
            // [request, response, case],
            ['{', '{"form":["Request cannot be parsed as JSON"]}', 'Broken JSON request'],
            ['{"datetime":["This value is not valid."]}', '{"datetime":["This value is not valid."]}', 'Invalid datetime'],
            ['{}', '{"datetime":["This value should not be blank."]}', 'Datetime missing'],
        ];
    }

    private function sendRequest(KernelBrowser $client, string $content = null)
    {
        $client->request(
            Request::METHOD_POST,
            '/v1/earth2mars',
            [],
            [],
            [
                'HTTP_CONTENT_TYPE' => 'application/json',
                'PHP_AUTH_USER' => 'test',
                'PHP_AUTH_PW'   => 'test',
            ],
            $content
        );
    }

}