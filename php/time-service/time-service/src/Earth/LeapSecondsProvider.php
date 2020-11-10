<?php

namespace App\Earth;

/**
 * @link https://en.wikipedia.org/wiki/Leap_second
 */
class LeapSecondsProvider
{
    public function getCurrent(): int
    {
        // TODO: can be definitely better, but out of scope of this test task. I will create issue for that.
        return 37; // Sad that not 42 ;)
    }
}
