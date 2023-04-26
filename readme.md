error handling in golang:

In some functions, instead of one output, we can return multiple outputs.
Probably the second output is an object named error.
error is always the last parameter that is returned and needs to be handled.
Here are 4 strategies for handling error.

1-error propagation:
When an error occurs in the lower coding layers, we must bring this error to the upper layers and add details to it so that the person receiving it can see exactly what happened in detail.
In the example given in the code, the anonymous function is responsible for making a request to the written URL, since this is a virtual server, it gives an error.
The function also returns an error in the output, but by handling the error, we made this error to be sent with a specific message.

2-retry on error:
The error may be a temporary error, for example, we want to make a connection to the database and at the same moment the database does not give us a connection (for example, because the connections are full at that time)
  These problems are not permanent and may be fixed moments later.
Other examples can be mentioned in these two cases:
error 500 in API call
network error in API call

In these cases, it is better to perform a series of retry operations on this error before the error function returns.
In the function we wrote for this strategy, we defined a time of 15 seconds using the deadline, then with the for loop we say that you have a time period of 15 seconds to do the things that your card here is http get request.
If this room does not occur, sleep as an exponential function.
In order not to engage the server too much, we must create a suitable time interval.
Exponential function:
The first time zero seconds, the second time one second, the fifteenth time, sixteen seconds
Then continue.
If there is no error, return the resp and consider the error as filled.
In the deadline is finished section, we said that if the deadline is over and you are not successful in doing the work, display this text with an error.

3-log and exit
For example, let's assume that API call is an important and vital task, if it is not done, the rest of the work will be disrupted.
In fact, if this task is not done correctly, the whole program is useless, and if there is an error in this part, we have to exit the program.
We can define to use log .fatal if there is an error.
The log .fatal command can close the program.
The important point in all these strategies is our client's main function.

4-log and continue :
Here, the task we have is not as critical as the third case and only involves or disables limited features of the program.
In this case, we write a log and return a nil instead of what was the output of the program (in this example, a pointer).

