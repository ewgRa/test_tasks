<?php

namespace App\Mars;

use App\Earth\LeapSecondsProvider;

/**
 * Developer experience from this class should be similar to php \DateTime class.
 * Probably there is something to think on it. For now it is quick and dirty idea.
 * Also as "new Msd2MtcFormatter()" doesn't look super nice.
 */
class MarsDateTime
{
    private float $msd;

    public static function createFromMsd(float $msd): self
    {
        $instance = new self();
        $instance->msd = $msd;

        return $instance;
    }

    /**
     * @throws \Exception if format is not supported
     */
    public function format($format): string
    {
        switch ($format) {
            case 'S': // return MSD
                return (string)$this->msd;
            case 'H:i:s': // return MTC
                $formatter = new Msd2MtcFormatter();
                return (string)$formatter->format($this->msd, $format);
            default:
                throw new \Exception("Format is not supported, sorry, maybe later");
        }
    }
}
