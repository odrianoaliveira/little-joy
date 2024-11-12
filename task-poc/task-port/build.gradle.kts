import com.google.protobuf.gradle.id

plugins {
    java
    alias(libs.plugins.protobuf)
}

dependencies {
    implementation("jakarta.annotation:jakarta.annotation-api:1.3.5")
    implementation("com.google.protobuf:protobuf-java:4.28.3")
    implementation("io.grpc:grpc-protobuf:1.68.1")
    implementation("io.grpc:grpc-stub:1.68.1")

    testImplementation("org.junit.jupiter:junit-jupiter:5.11.3")
    testImplementation("org.mockito:mockito-core:5.14.2")
    testImplementation("io.grpc:grpc-testing:1.68.1")
}

protobuf {
    protoc {
        artifact = libs.protoc.asProvider().get().toString()
    }
    plugins {
        id("grpc") {
            artifact = libs.protoc.gen.grpc.java.get().toString()
        }
    }
    generateProtoTasks {
        all().forEach {
            it.plugins {
                create("grpc")
            }
        }
    }
}

java {
    withJavadocJar()
    withSourcesJar()
}

tasks.jar {
    archiveBaseName.set("task-port")
    archiveVersion.set("1.0.0")
    from(sourceSets.main.get().output)
}

tasks.named<Test>("test") {
    useJUnitPlatform()
}