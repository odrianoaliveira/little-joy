package task.port;

import java.util.List;

import io.grpc.stub.StreamObserver;

public abstract class TaskPort<T> extends TaskServiceGrpc.TaskServiceImplBase {

  private final TaskMapper<Task, T> taskMapper;

  public TaskPort(TaskMapper<Task, T> taskMapper) {
    this.taskMapper = taskMapper;
  }

  @Override
  public void getTask(TaskRequest request, StreamObserver<TaskResponse> responseObserver) {
    var tasks = listTasksFromAdapter().stream()
        .map(taskMapper::mapTo)
        .toList();
    var page = Page.newBuilder()
        .setCurrent(0)
        .setSize(10)
        .setTotal(1)
        .build();
    var response = TaskResponse.newBuilder()
        .addAllTasks(tasks)
        .setPage(page)
        .build();
    responseObserver.onNext(response);
    responseObserver.onCompleted();
  }

  abstract List<T> listTasksFromAdapter();
}
