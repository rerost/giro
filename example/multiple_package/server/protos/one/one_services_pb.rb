# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: protos/one/one.proto for package 'example.multiple_package.protos.one'

require 'grpc'
require 'protos/one/one_pb'

module Example
  module MultiplePackage
    module Protos
      module One
        module GiroService
          class Service

            include ::GRPC::GenericService

            self.marshal_class_method = :encode
            self.unmarshal_class_method = :decode
            self.service_name = 'example.multiple_package.protos.one.GiroService'

            rpc :GiroTest1, ::Example::MultiplePackage::Protos::One::GiroTestRequest1, ::Example::MultiplePackage::Protos::One::GiroTestResponse1
            rpc :GiroTest2, ::Example::MultiplePackage::Protos::One::GiroTestRequest2, ::Example::MultiplePackage::Protos::One::GiroTestResponse2
            rpc :GiroEmptyTest, ::Google::Protobuf::Empty, ::Google::Protobuf::Empty
          end

          Stub = Service.rpc_stub_class
        end
      end
    end
  end
end
