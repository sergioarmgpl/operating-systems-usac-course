import threading
import time

def task(name):
    print(f"Starting {name}")
    for i in range(0,5):
       time.sleep(0.5)
       print(f"{name}:"+str(i))
    print(f"Finished {name}")

# Create and start a thread
thread1 = threading.Thread(target=task, args=("Thread-1",))
thread2 = threading.Thread(target=task, args=("Thread-2",))
thread3 = threading.Thread(target=task, args=("Thread-3",))
thread1.start()
thread2.start()
thread3.start()

# Wait for the thread to complete
thread1.join()
print("join thread1")
thread2.join()
print("join thread2")
thread3.join()
print("join thread3")
print("All tasks done.")
