this_dir = File.expand_path(File.dirname(__FILE__))
$LOAD_PATH.unshift(this_dir) unless $LOAD_PATH.include?(this_dir)

require "protos/one/one_pb"
require "protos/two/two_pb"
require "protos/one/one_services_pb"
require "protos/two/two_services_pb"

require 'grpc'

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

def main
  s = GRPC::RpcServer.new
  addr = "0.0.0.0:5001"
  s.add_http2_port(addr, :this_port_is_insecure)
  puts "Listing " + addr

  s.handle(GiroService.new)
  s.handle(BqvService.new)
  s.run_till_terminated
end

main
