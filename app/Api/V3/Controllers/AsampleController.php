<?php

namespace App\Api\V3\Controllers;

class AsampleController extends Controller
{
    public function __invoke()
    {
        return response()->json("hello");
    }
}
