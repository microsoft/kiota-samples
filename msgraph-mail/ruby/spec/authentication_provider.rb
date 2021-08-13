require 'concurrent'

class AuthenticationProvider
  include Concurrent::Async

  def get_authorization_token(request_url)
    return '' # Throw Token here or replace with auth Provider
  end

end
