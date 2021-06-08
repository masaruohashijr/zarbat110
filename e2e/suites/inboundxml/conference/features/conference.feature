Feature: conference    
  In order to make a conference call
  As an end user
  I want to dial to a number, create and set a conference room

  Scenario: Make a Conference // WE CAN USE OUTLINE and BACKGROUND HERE

    Given I have my "Account" 
    And I set "+333" to say "What a wonderful world"
    And I set "+444" to say "The colors of the rainbow"
    When I make a call from "+111" to a "+222" and set as a conference room with size 3
    And  I make a call from "+333" to a "+222" 
    And  I make a call from "+444" to a "+222" 
    Then I should get the text "What a wonderful world" from the number "+333"
    And I should get the text "The colors of the rainbow" from the number "+444"

