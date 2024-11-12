plugins {
    application
    alias(libs.plugins.kotlin.jvm)
}

dependencies {
    implementation(project(":task-port"))
    implementation("io.grpc:grpc-kotlin-stub:1.3.0")
    implementation("io.grpc:grpc-protobuf:1.42.1")
}

application {
    mainClass.set("task.adapter.poc.MainKt")
}

tasks.named<Test>("test") {
    useJUnitPlatform()
}
