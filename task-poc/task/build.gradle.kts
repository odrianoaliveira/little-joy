
plugins {
    application
}

dependencies {
    testImplementation("org.jetbrains.kotlin:kotlin-test-junit5")
    testImplementation("org.junit.jupiter:junit-jupiter-engine:5.9.1")

    implementation("com.google.guava:guava:31.1-jre")
}

application {
    mainClass.set("task.adapter.poc.MainKt")
}

tasks.named<Test>("test") {
    useJUnitPlatform()
}
