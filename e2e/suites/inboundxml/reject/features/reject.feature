Feature: reject    
  In order to reject a call
  As an end user
  I want to try to make a call to a Number and listen to a not-in-service message.

  Scenario: Pause a sequence of sentences

    Given I have my "Account"
    And I set "+888" to dial to "+999"
    And gather speech within 5 seconds
    And I set "+999" to reject calls with "not-in-service" reason
    When I try to make a call from "+888" to "+999"    
    Then I get a "not-in-service" reason in plain text