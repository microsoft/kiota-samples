<?php

use GetUser\Client\GraphApiClient;
use Microsoft\Kiota\Abstractions\ApiException;
use Microsoft\Kiota\Authentication\Oauth\ClientCredentialContext;
use Microsoft\Kiota\Authentication\PhpLeagueAuthenticationProvider;
use Microsoft\Kiota\Http\GuzzleRequestAdapter;

require __DIR__.'/vendor/autoload.php';

try {
    $tenantId = 'tenantId';
    $clientId = 'clientId';
    $clientSecret = 'secret';
    $userId = 'userPrincipalName';

    $allowedHosts = ['graph.microsoft.com'];
    $scopes = ['https://graph.microsoft.com/.default'];

    $tokenRequestContext = new ClientCredentialContext(
        $tenantId,
        $clientId,
        $clientSecret
    );

    $authProvider = new PhpLeagueAuthenticationProvider($tokenRequestContext, $scopes, $allowedHosts);
    $requestAdapter = new GuzzleRequestAdapter($authProvider);
    $client = new GraphApiClient($requestAdapter);

    $user = $client->usersById($userId)->get()->wait();
    echo "Hello {$user->getDisplayName()}, your ID is {$user->getId()}";

} catch (ApiException $ex) {
    echo $ex->getMessage();
}
