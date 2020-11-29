<?php

namespace App\Tests\Mars;

use App\Mars\Msd2MtcFormatter;
use PHPUnit\Framework\TestCase;

class Msd2MtcFormatterTest extends TestCase
{
    /**
     * @dataProvider cases
     */
    public function testConvert(float $msd, string $mtcExpected)
    {
        $formatter = new Msd2MtcFormatter();

        $this->assertSame($mtcExpected, $formatter->format($msd, 'H:i:s'));
    }

    public function cases(): array
    {
        return [
            // [msd, mtcExpected],
            [87450.37758147829, '09:03:43'],
            [52020.09484639508, '02:16:34'],
        ];
    }
}
