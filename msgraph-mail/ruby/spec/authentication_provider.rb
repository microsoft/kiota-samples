class AuthenticationProvider < MicrosoftKiotaAbstractions::BaseBearerTokenAuthenticationProvider
  def get_authorization_token(request)
    return "" # TODO replace with access token
  end
end
