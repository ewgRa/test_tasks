<?php

namespace App\Tests\Mars;

use App\Earth\LeapSecondsProvider;
use App\Mars\Earth2MsdConverter;
use PHPUnit\Framework\TestCase;

class Earth2MsdConverterTest extends TestCase
{
    /**
     * @dataProvider cases
     */
    public function testConvert(string $datetime, float $msdExpected)
    {
        $converter = new Earth2MsdConverter($this->leapSecondsProviderMock());
        $this->assertSame($msdExpected, $converter->convert(new \DateTime($datetime)));
    }

    public function cases(): array
    {
        return [
            // [datetime, msdExpected],
            ['2120-01-04T01:00:00+0100', 87450.37758147829],
            ['2020-05-02T16:40:00+0000', 52020.09484639508],
        ];
    }

    private function leapSecondsProviderMock(): LeapSecondsProvider
    {
        $mock = $this->createMock(LeapSecondsProvider::class);

        $mock->expects($this->once())
            ->method('getCurrent')
            ->will($this->returnValue(37))
        ;

        /** @var LeapSecondsProvider $mock */
        return $mock;
    }
}
