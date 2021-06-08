Feature: record    
  In order to record a call
  As an end user
  I want to try to make a call and record this call until key "#" is pressed.

  Scenario: Pause a sequence of sentences

    Given I have my "Account"
    And I set "+888" a "woman" to say "You would never break the chain." with a 5 seconds timeout
    And I set "+999" to record 5 seconds of the call
    When I make a call from "+888" to "+999"    
    Then I get the transcription "You would never break the chain."