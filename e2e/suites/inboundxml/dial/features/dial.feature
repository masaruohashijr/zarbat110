Feature: dial    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call

  Scenario: Make a Transfer

    Given I have my "Account"
    And I set "+999" to dial to "+888"
    And I set "+888" to say "I see skies of blue"
    When I make a call to the number "+999"
    Then this call is transfered to the number "+888"
    And I get the response "I see skies of blue" in plain text