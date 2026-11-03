require 'webmock/rspec'

RSpec.describe 'main script' do
  it 'calls the posts API' do
    posts_url = 'https://jsonplaceholder.typicode.com/posts'
    post_json = '{ "userId": 1, "id": 5, "title": "Post title", "body": "Post body" }'
    json_headers = { content_type: 'application/json' }

    stub_request(:get, posts_url)
      .to_return({ body: "[#{post_json}, #{post_json}]", headers: json_headers, status: 200 })
    stub_request(:get, "#{posts_url}/5")
      .to_return({ body: post_json, headers: json_headers, status: 200 })
    stub_request(:post, posts_url)
      .to_return({ body: '{ "userId": 42, "id": 101 }', headers: json_headers, status: 201 })
    stub_request(:patch, "#{posts_url}/5")
      .to_return({ body: '{ "userId": 1, "id": 5, "title": "Updated title", "body": "Post body" }',
                   headers: json_headers, status: 200 })
    stub_request(:delete, "#{posts_url}/5")
      .to_return({ body: '{}', headers: json_headers, status: 200 })

    iseq = RubyVM::InstructionSequence.compile_file './main.rb'
    expect { iseq.eval }.to output(
      "Retrieved 2 posts.\n" \
      "Retrieved post - ID: 5, Title: Post title, Body: Post body\n" \
      "Created new post with ID: 101\n" \
      "Updated post - ID: 5, Title: Updated title, Body: Post body\n"
    ).to_stdout

    expect(a_request(:post, posts_url)
      .with(body: { userId: 42, title: 'Testing Kiota-generated API client', body: 'Hello world!' }))
      .to have_been_made
    expect(a_request(:patch, "#{posts_url}/5").with(body: { title: 'Updated title' })).to have_been_made
    expect(a_request(:delete, "#{posts_url}/5")).to have_been_made
  end
end
