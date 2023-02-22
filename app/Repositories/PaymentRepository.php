<?php

namespace App\Repositories;

interface PaymentRepository
{
    public function chargeCard(array $data);
}
