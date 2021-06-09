Feature: play    
  In order to play a song when someone calls my number
  As an end user
  I want to set a MP3 to play every time another number calls mine.

  Scenario: Play a MP3

    Given my test setup runs #-> myTestSetupRuns()
    And "NumberB" configured to play "music.mp3" #-> configuredToPlay(numberB string, music string)
    When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)
    Then "NumberA" should get track name "trackName" #-> shouldGetTrackName(numberA string, trackName string)    