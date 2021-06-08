Feature: hangup    
  In order to hangup a call after a certain time in seconds
  As an end user
  I want to have my call automatically off

  Scenario: Hangup a call

    Given I have my "Account"
    And I set "+888" to record the call
    And I set "+999" to hangup after 3 seconds
    When I make a call from "+888" to "+999"
    And 3 seconds has past 
    Then the call should be off
    And I can see that the recording lasted 3 seconds 