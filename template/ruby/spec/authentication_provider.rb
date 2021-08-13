require 'concurrent'

class AuthenticationProvider
  include Concurrent::Async

  def get_authorization_token(request_url)
    return ''
  end

end
