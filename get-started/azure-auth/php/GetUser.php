<!-- Copyright (c) Microsoft Corporation.
     Licensed under the MIT License.-->

<!-- <ProgramSnippet> -->
<?php

use GetUser\Client\GraphApiClient;
use Microsoft\Kiota\Abstractions\ApiException;
use Microsoft\Kiota\Authentication\Oauth\AuthorizationCodeContext;
use Microsoft\Kiota\Authentication\PhpLeagueAuthenticationProvider;
use Microsoft\Kiota\Http\GuzzleRequestAdapter;

require __DIR__.'/vendor/autoload.php';

try {
    $clientId = 'clientId';
    $clientSecret = 'secret';
    $authorizationCode = 'authCode';

    $tenantId = 'common';
    $redirectUri = 'http://localhost';

    $scopes = ['User.Read'];

    $tokenRequestContext = new AuthorizationCodeContext(
        $tenantId,
        $clientId,
        $clientSecret,
        $authorizationCode,
        $redirectUri
    );

    $authProvider = new PhpLeagueAuthenticationProvider($tokenRequestContext, $scopes);
    $requestAdapter = new GuzzleRequestAdapter($authProvider);
    $client = new GraphApiClient($requestAdapter);

    $me = $client->me()->get()->wait();
    echo "Hello {$me->getDisplayName()}, your ID is {$me->getId()}";

} catch (ApiException $ex) {
    echo $ex->getMessage();
}
?>
<!-- </ProgramSnippet> -->
