package task.port;

import java.util.List;

import io.grpc.stub.StreamObserver;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;

class TaskPortTest {

  @Test
  void shouldListAndMapTasks() {
    var port = new TaskPort<>((TaskMapper<Task, TaskDto>) dto -> Task.newBuilder()
        .setId(dto.id)
        .setTitle(dto.title)
        .setStatus(TaskStatus.OPEN)
        .setDescription("an external task")
        .build()) {

      @Override
      List<TaskDto> listTasksFromAdapter() {
        return List.of(new TaskDto("1", "task 1"));
      }
    };

    var request = TaskRequest.newBuilder().setId("1").build();
    var responseObserver = mock(StreamObserver.class);
    var responseCaptor = ArgumentCaptor.forClass(TaskResponse.class);

    port.getTask(request, responseObserver);
    verify(responseObserver).onNext(responseCaptor.capture());
    verify(responseObserver).onCompleted();

    var result = responseCaptor.getAllValues();
    assertEquals(1, result.size());
    var expectedTask = Task.newBuilder()
        .setId("1")
        .setTitle("task 1")
        .setStatus(TaskStatus.OPEN)
        .setDescription("an external task")
        .build();
    assertEquals(expectedTask, result.get(0).getTasksList().get(0));
  }

  record TaskDto(String id, String title) {

  }
}