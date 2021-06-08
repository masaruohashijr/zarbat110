Feature: play    
  In order to play a song when someone calls my number
  As an end user
  I want to set a MP3 to play every time another number calls mine.

  Scenario: Play a MP3

    Given I have my "Account"
    And I set "+999" to play "music.mp3"
    When I make a call from "+888" to "+999"
    Then I listen to "music.mp3"
    And I get "track", "author", "album" and "releaseYear"