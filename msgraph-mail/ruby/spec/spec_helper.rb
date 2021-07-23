# frozen_string_literal: true

require 'microsoft_kiota_abstractions'
require_relative './../graphrubyv4.rb'
# require_relative './../api_client'
# require_relative './../users/entity'
# require_relative './../users/attachment.rb'
# require_relative './../users/inference_classification/inference_classification_override'
# Dir['./../*.rb'].each { |file| require file }
# Dir['./../users/*.rb'].each { |file| require file }
# Dir['./../users/inference_classification/*.rb'].each { |file| require file }
# Dir['./../users/item/*.rb'].each { |file| require file }
# Dir['./../users/mail_folders/*.rb'].each { |file| require file }
# Dir['./../users/messages/*.rb'].each { |file| require file }
Dir["#{File.dirname(__FILE__)}/./../graphrubyv4/**/*.rb"].each { |f| load(f) }

RSpec.configure do |config|
  # Enable flags like --only-failures and --next-failure
  config.example_status_persistence_file_path = ".rspec_status"

  # Disable RSpec exposing methods globally on `Module` and `main`
  config.disable_monkey_patching!

  config.expect_with :rspec do |c|
    c.syntax = :expect
  end
end
