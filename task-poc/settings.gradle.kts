dependencyResolutionManagement {
    repositories {
        mavenCentral()
    }
}

rootProject.name = "task-adapter-poc"

include("task")
include("task-port")
include("task-adapter")
