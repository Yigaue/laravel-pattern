<?php

namespace App\Api\V3\Controllers;

use App\Repositories\PaymentRepository;
use Illuminate\Http\Request;

class PaymentController extends Controller
{
    public $paymentRepository;

    // Always program to an interface;

    public function __construct(PaymentRepository $paymentRepository)
    {
        $this->paymentRepository = $paymentRepository;
    }

    public function chargeCard()
    {
        $request = new Request([
            'amount' => 100.00,
            'card' => '283948374948379483748'
        ]);
        // This will be coming from a form  or API through Request injected in the method chargeCard

        $response =  $this->paymentRepository->chargeCard($request->all());

        return response()->json($response, $response['code']);
    }
}
