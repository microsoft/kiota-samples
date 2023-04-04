require_relative 'static_access_token_provider'
class AuthenticationProvider < MicrosoftKiotaAbstractions::BaseBearerTokenAuthenticationProvider
  def initialize()
    super(StaticAccessTokenProvider.new())
  end 
end
