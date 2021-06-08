Feature: calls    
  In order to list my account calls
  As a developer user
  I want to consume an API List Calls service 

  @Success
  Scenario: List Calls

    Given I have my "AccountSid" 
    And I have my "AuthToken"
    When I send "POST" request to "/Accounts/{AccountSID}/Calls.json"
    And I send a basic authentication with "AccountSid" and "AuthToken"
    Then the response code should be 200
    And the response should contains json:
      '''
      {
      "calls" : "calls"
      }
      ''' 
  
  @list-calls-error
  Scenario: List Calls without AccountSid

    Given I dont have my "AccountSid" 
    And I have my "AuthToken"
    When I send "POST" request to "/Accounts/{AccountSID}/Calls.json"
    And I send a basic authentication with "AccountSid" and "AuthToken"
    Then the response code should be 401
    And the response should be:
      """
      {"error":"Unauthorized"}
      """ 

  @list-calls-error
  Scenario: List Calls without AuthToken

    Given I have my "AccountSid" 
    And I dont have my "AuthToken"
    When I send "POST" request to "/Accounts/{AccountSID}/Calls.json"
    And I send a basic authentication with "AccountSid" and "AuthToken"
    Then the response code should be 401
    And the response should be:
      """
      {"error":"Unauthorized"}
      """ 