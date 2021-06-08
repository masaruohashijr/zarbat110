Feature: play last recording    
  In order to listen the last recording of my calls
  As an end user
  I want to call a number and set to play the last recording.

  Scenario: Play Last Recording

    Given I have my "Account"
    And I set "+666" to gather speech within 5 seconds of timeout
    And I set "+777" a "woman" to say "This is the last recording"
    And I set "+888" to record the call with a 5 seconds timeout
    And I set "+999" to play my last recording
    When I make a call from "+666" to "+999"    
    Then I get the transcription "This is the last recording"