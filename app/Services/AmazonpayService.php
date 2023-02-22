<?php

namespace App\Services;

use App\Repositories\PaymentRepository;

class AmazonpayService implements PaymentRepository
{
    /**
     * Always accept and return array in method
     *
     * @param array $data
     * @return array
     */
    public function chargeCard(array $data)
    {
        $card = $data['card'];
        $amount = $data['amount'];
        return [
            'status' => 'success',
            'message' => sprintf('your card :%s has been charged %s USD', $card, $amount),
            'data' => $data,
            'code' => 200
        ];
    }
}
