# frozen_string_literal: true

module GetUser
  # Loads every model and enum registered for autoload, for example before a server forks.
  def self.eager_load!
    loop do
      loaded = false
      pending = [self]
      seen = {}
      until pending.empty?
        mod = pending.pop
        next if seen.key?(mod)

        seen[mod] = true
        mod.constants(false).each do |name|
          loaded = true if mod.autoload?(name, false)
          value = mod.const_get(name, false)
          pending << value if value.instance_of?(::Module)
        end
      end
      break unless loaded
    end
  end
end
