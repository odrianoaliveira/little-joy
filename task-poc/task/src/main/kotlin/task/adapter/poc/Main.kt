package task.adapter.poc

class Main {
    val greeting: String
        get() {
            return "Hello World!"
        }
}

fun main() {
    println(Main().greeting)
}
