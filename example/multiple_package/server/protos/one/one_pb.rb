# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/one/one.proto

require 'google/protobuf'

descriptor_data = "\n\x14protos/one/one.proto\x12#example.multiple_package.protos.one\"#\n\x10GiroTestRequest1\x12\x0f\n\x07message\x18\x01 \x01(\t\"\x12\n\x10GiroTestRequest2\"$\n\x11GiroTestResponse1\x12\x0f\n\x07message\x18\x01 \x01(\t\"\x13\n\x11GiroTestResponse22\x89\x02\n\x0bGiroService\x12|\n\tGiroTest1\x12\x35.example.multiple_package.protos.one.GiroTestRequest1\x1a\x36.example.multiple_package.protos.one.GiroTestResponse1\"\x00\x12|\n\tGiroTest2\x12\x35.example.multiple_package.protos.one.GiroTestRequest2\x1a\x36.example.multiple_package.protos.one.GiroTestResponse2\"\x00\x42\x43ZAgithub.com/rerost/giro/example/multiple_package/protos/one;one_pbb\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool
pool.add_serialized_file(descriptor_data)

module Example
  module MultiplePackage
    module Protos
      module One
        GiroTestRequest1 = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("example.multiple_package.protos.one.GiroTestRequest1").msgclass
        GiroTestRequest2 = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("example.multiple_package.protos.one.GiroTestRequest2").msgclass
        GiroTestResponse1 = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("example.multiple_package.protos.one.GiroTestResponse1").msgclass
        GiroTestResponse2 = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("example.multiple_package.protos.one.GiroTestResponse2").msgclass
      end
    end
  end
end
