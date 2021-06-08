Feature: say    
  In order to read text to the Number B (called) using a text-to-speech engine   
  As an end user
  I want that Number A (caller) listen the speech set to be read on Number B.

  Scenario: Say something

    Given I have my "Account" 
    And I set "+888" to say "I think to myself"
    And I set "+999" to gather speech result within 15 seconds
    When I make a call from "+888" to "+999"
    Then I should listen "I think to myself"