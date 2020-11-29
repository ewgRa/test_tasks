<?php

namespace App\Mars;

use App\Earth\LeapSecondsProvider;

/**
 * @link https://www.giss.nasa.gov/tools/mars24/help/algorithm.html
 * @link https://msl-curiosity.requio.com/
 * @link https://jtauber.github.io/mars-clock/
 * TODO: what about dates before 1972? Maybe later, I will create an issue :(
 */
class Earth2MsdConverter
{
    private const TAI_UTC_CORRECTION = 32.184;
    private const ONE_DAY_SECONDS = 86400;
    private const EPOCH_JULIAN_SECONDS = 2440587.5 * self::ONE_DAY_SECONDS;
    private const JULIAN_2000_EPOCH = 2451549.5 * self::ONE_DAY_SECONDS;
    private const MIDNIGHT_ADJUSTMENT = 0.0009626 * self::ONE_DAY_SECONDS;
    private const MSD_POSITIVE_OFFSET = 44796 * self::ONE_DAY_SECONDS;
    private const MARTIAN_EARTH_DAY_RATIO = 1.0274912517;

    private LeapSecondsProvider $leapSecondsProvider;

    public function __construct(LeapSecondsProvider $leapSecondsProvider)
    {
        $this->leapSecondsProvider = $leapSecondsProvider;
    }

    public function convert(\DateTime $dateTime): float
    {
        $jdut = $this->julian($dateTime->getTimestamp());
        $jdtt = $this->terrestrialTime($jdut);
        $deltaT = $jdtt - self::JULIAN_2000_EPOCH;
        $msdSeconds = $deltaT / self::MARTIAN_EARTH_DAY_RATIO + self::MSD_POSITIVE_OFFSET - self::MIDNIGHT_ADJUSTMENT;

        return $msdSeconds / self::ONE_DAY_SECONDS;
    }

    private function julian(int $timestamp): float
    {
        return self::EPOCH_JULIAN_SECONDS + $timestamp;
    }

    private function terrestrialTime(float $days): float
    {
        $terrestrialCorrection = ($this->leapSecondsProvider->getCurrent() + self::TAI_UTC_CORRECTION);

        return $days + $terrestrialCorrection;
    }
}
