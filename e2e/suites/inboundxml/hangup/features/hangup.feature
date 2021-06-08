Feature: hangup    
  In order to hangup a call after a certain time in seconds
  As an end user
  I want to have my call automatically off

  Scenario: Hangup a call

    Given I have my "Account"
    When I make a call to "999"
    And after 3 seconds past 
    Then the call should be off