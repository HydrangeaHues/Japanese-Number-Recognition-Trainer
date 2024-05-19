Functionality Summary

I study Japanese but struggle with being able to parse and understand numbers when said in conversation. This small program effectively runs a quiz by generating a random number between 1 - 10000
and then uses TTS to speak that number in Japanese. The user has 10 seconds from the start of the number being spoken to enter their answer for what number was said or enter -1 to have the number
be replayed. At the end of the 10 seconds, the quiz will move onto the next question. The quiz is currently locked to 5 questions. At the end of the quiz, the number of correct answers is reported.


Requirements

mplayer - The TTS package I used (htgotts) gives various options for handlers when playing the MP3 files created for TTS. I attempted to use the Native handler but was having an issue where
          subsequent MP3 files would not be played after playing the first. This issue was resolved by using the mplayer handler, so you will need to install mplayer before being able to use this program.


Future Improvement Ideas

1. Clearing and recreating the audio file on every quiz question is not ideal. Finding a more efficient way / point in time to do this is an area of improvement.
2. I also struggle with date and time recognition, so I could abstract out the functions for playing audio, polling for answers, and running the quiz and then add code to randomly generate a date or time. From there I could allow users to choose the type of quiz they want.
3. Allowing users to choose the number of questions, the timeout limit, and potentially even the playback language would be a relatively small adds that would give users more agency over their experience.


Ultimately this is just a fun tool to help me with my language learning. Feel free to use it if it helps you too.
