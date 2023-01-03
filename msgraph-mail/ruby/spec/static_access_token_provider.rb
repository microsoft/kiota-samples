class StaticAccessTokenProvider
    include MicrosoftKiotaAbstractions::AccessTokenProvider
    def get_authorization_token(request, additional_properties = {})
      Fiber.new do
        "" # TODO replace with access token
      end
    end
end
  