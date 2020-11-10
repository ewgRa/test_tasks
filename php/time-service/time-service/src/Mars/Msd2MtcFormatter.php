<?php

namespace App\Mars;

class Msd2MtcFormatter
{
    private const SECONDS_IN_HOUR = 3600;
    private const HOURS_IN_DAY = 24;

    public function format(float $msd, string $format): string
    {
        $hours = fmod(self::HOURS_IN_DAY * $msd, self::HOURS_IN_DAY);
        $seconds = (int)($hours * self::SECONDS_IN_HOUR);

        return \DateTime::createFromFormat('U', $seconds)->format($format);
    }
}
