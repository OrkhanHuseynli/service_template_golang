@service
Feature: As a user when I call the service endpoint for "product", I would like to receive a json response with
  a corresponding message

  Scenario: Invalid query
    Given I have no product criteria
    When I call the product endpoint
    Then I should receive a bad request message

  Scenario: Valid query
    Given I have a valid product criteria
    When I call the product endpoint
    Then I should receive a json response with a corresponding message