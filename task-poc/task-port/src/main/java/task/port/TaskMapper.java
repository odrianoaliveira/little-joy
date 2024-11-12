package task.port;

public interface TaskMapper<Task, T> {

  task.port.Task mapTo(T dto);
}
