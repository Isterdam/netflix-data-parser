# Netflix Data Parser

I wrote this simple program as a part of a project in a course on networking (EITF45 at LTH). My lab partner and I were to examine an online service that interested us, and we picked Netflix. 
More precisely, we wanted to examine how Netflix buffered their streams with regards to video quality (Netflix utilizes MPEG-DASH with a bit rate adaptation algorithm), and how the buffer behaved for different network speeds. To enable
this, I was firstly able to limit my network speed in the router settings (we tested for 0.5, 1.5 and 5 Mbps downstream). Secondly, I wrote 
a script that copied the data from Netflix's built-in diagnostic tool which could ultimately be downloaded as a pdf in the browser. This pdf
was fed to the Golang program, which ultimately picked out the measurements of interest (playing bit rate, buffer bit rate and buffer length) 
and automatically visualized them in a graph. These were the results:

<p float="left" align="center">
  <img src="https://github.com/Isterdam/netflix-data-parser/blob/master/project/0.5.png" width="300" />
  &nbsp;&nbsp;
  &nbsp;&nbsp;
  &nbsp;&nbsp;
  &nbsp;&nbsp;
  &nbsp;&nbsp;
  <img src="https://github.com/Isterdam/netflix-data-parser/blob/master/project/1.5.png" width="300" /> 
</p>

<p align="center">
<img src="https://github.com/Isterdam/netflix-data-parser/blob/master/project/5.png" width="300" />
</p>

The full report is available [here](https://github.com/Isterdam/netflix-data-parser/blob/master/project/EITF45_Projekt.pdf), albeit in Swedish.
