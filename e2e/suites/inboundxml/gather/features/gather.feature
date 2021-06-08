Feature: gather    
  In order to input digits and a speech that the call should listen for
  As an end user
  I want to gather the input digits or the speech said.

  Scenario: Gather DTMF

    Given I have my "Account"  
    And I set "+888" to send digits "123"
    When I make a call from "+888" to "+999" 
    Then I should be able to see "123" in plain text

  Scenario: Gather Speech

    Given I have my "Account"  
    And I set "+888" to say "something"
    When I make a call to "+999" 
    Then I should be able to see "something" in plain text