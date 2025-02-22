this_dir = File.expand_path(File.dirname(__FILE__))
$LOAD_PATH.unshift(this_dir) unless $LOAD_PATH.include?(this_dir)

require "protos/one/one_pb"
require "protos/two/two_pb"
require "protos/one/one_services_pb"
require "protos/two/two_services_pb"

require 'grpc'
require 'grpc/health/v1/health_services_pb'
require 'grpc/reflection/v1alpha/reflection'

class GiroService < Example::MultiplePackage::Protos::One::GiroService::Service
  def giro_test1(req, _call)
    return Example::MultiplePackage::Protos::One::GiroTestResponse1.new(message: req.message)
  end

  def giro_test2(_req, _call)
    return Example::MultiplePackage::Protos::One::GiroTestResponse2.new
  end
end

class BqvService < Example::MultiplePackage::Protos::Two::BqvService::Service
  def bqv_test1(_req, _call)
    return Example::MultiplePackage::Protos::Two::BqvTestResponse1.new
  end

  def bqv_test2(_req, _call)
    return Example::MultiplePackage::Protos::Two::BqvTestResponse2.new
  end
end

class HealthCheckService < Grpc::Health::V1::Health::Service
  def check(health_check_request, _unused_call)
    puts "Health check called for service: #{health_check_request&.service || 'unknown'}"
    Grpc::Health::V1::HealthCheckResponse.new(
      status: Grpc::Health::V1::HealthCheckResponse::ServingStatus::SERVING
    )
  end
end

def main
  s = GRPC::RpcServer.new
  port = ENV.fetch('APP_PORT', '5001')
  addr = "0.0.0.0:#{port}"
  s.add_http2_port(addr, :this_port_is_insecure)
  puts "Starting gRPC server on #{addr}"
  puts "Registering services..."

  s.handle(GiroService.new)
  s.handle(BqvService.new)
  s.handle(HealthCheckService.new)
  s.handle(GRPC::Reflection::V1alpha::ServerReflection::Service.new)
  s.run_till_terminated
end

main
