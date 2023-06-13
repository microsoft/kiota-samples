require 'webmock/rspec'

RSpec.describe 'get_user script' do
  it 'compiles' do
    stub_request(:post, 'https://login.microsoftonline.com/common/oauth2/v2.0/token')
      .to_return({
        body: '{ "access_token": "token", "token_type": "Bearer", "expires_in": 3599, "scope": "scope", "refresh_token": "refresh" }',
        headers: {content_type: 'application/json'},
        status: 200
      })

    stub_request(:get, 'https://graph.microsoft.com/v1.0/me')
      .to_return({
        body: '{ "displayName": "Adele Vance", "id": "user_id" }',
        headers: {content_type: 'application/json'},
        status: 200
      })
    iseq = RubyVM::InstructionSequence.compile_file './get_user.rb'
    output = iseq.eval
  end
end
