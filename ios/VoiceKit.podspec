Pod::Spec.new do |s|
  s.name     = "VoiceKit"
  s.version  = "1.0.0"
  s.license  = "Apache License, Version 2.0"
  s.authors  = { 'Tinkoff VoiceKit Team' => 'speech@tinkoff.ru' }
  s.homepage = "https://voicekit.tinkoff.ru"
  s.summary = "VocieKit gRPC APIs"
  s.source = { :git => 'https://github.com/TinkoffCreditSystems/voicekit-examples.git' }
  s.platform = :ios

  s.ios.deployment_target = "10.0"
  s.swift_version = '5.0'

  # Base directory where the .proto files are.
  src = "../apis"
  googleapis = "../third_party/googleapis"

  # Need at least gRPC >= 1.21 due to networking transition issues when using raw TCP
  # sockets as opposed to Apple's CFStream on gRPC < 1.21
  # See https://github.com/grpc/grpc/blob/master/src/objective-c/README-CFSTREAM.md
  # and https://github.com/grpc/grpc/blob/v1.19.0/src/objective-c/NetworkTransitionBehavior.md
  s.dependency '!ProtoCompiler-gRPCPlugin', '~> 1.21'

  # Pods directory corresponding to this app's Podfile, relative to the location of this podspec.
  pods_root = 'Pods'

  # Path where Cocoapods downloads protoc and the gRPC plugin.
  protoc_dir = "#{pods_root}/!ProtoCompiler"
  protoc = "#{protoc_dir}/protoc"
  plugin = "#{pods_root}/!ProtoCompiler-gRPCPlugin/grpc_objective_c_plugin"

  # Directory where the generated files will be placed.
  dir = "#{pods_root}/#{s.name}"

  # Run protoc with the Objective-C and gRPC plugins to generate protocol messages and gRPC clients.
  s.prepare_command = <<-CMD
    mkdir -p #{dir}
    #{protoc} \
        --plugin=protoc-gen-grpc=#{plugin} \
        --objc_out=#{dir} \
        --grpc_out=#{dir} \
        -I #{src} \
        -I #{googleapis} \
        -I #{protoc_dir} \
        #{src}/tinkoff/cloud/stt/v1/stt.proto \
        #{src}/tinkoff/cloud/tts/v1/tts.proto \
        #{googleapis}/google/api/annotations.proto \
        #{googleapis}/google/api/http.proto \
        #{protoc_dir}/google/protobuf/descriptor.proto
  CMD

  # Files generated by protoc
  s.subspec "Messages" do |ms|
    ms.source_files = "#{dir}/*.pbobjc.{h,m}", "#{dir}/**/*.pbobjc.{h,m}"
    ms.header_mappings_dir = dir
    ms.requires_arc = false
    # The generated files depend on the protobuf runtime.
    ms.dependency "Protobuf"
  end

  # Files generated by the gRPC plugin
  s.subspec "Services" do |ss|
    ss.source_files = "#{dir}/*.pbrpc.{h,m}", "#{dir}/**/*.pbrpc.{h,m}"
    ss.header_mappings_dir = dir
    ss.requires_arc = true
    # The generated files depend on the gRPC runtime, and on the files generated by protoc.
    ss.dependency "gRPC-ProtoRPC", "~> 1.21"
    ss.dependency "#{s.name}/Messages"
  end

  s.pod_target_xcconfig = {
    'GCC_PREPROCESSOR_DEFINITIONS' => '$(inherited) GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS=1',
    'DEFINES_MODULE' => 'YES',
    'CLANG_ALLOW_NON_MODULAR_INCLUDES_IN_FRAMEWORK_MODULES' => 'YES',
    'USER_HEADER_SEARCH_PATHS' => '"${PODS_ROOT}/VoiceKit"',
  }
end