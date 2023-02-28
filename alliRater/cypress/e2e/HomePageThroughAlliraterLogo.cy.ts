describe('Homepage through logo Test', () => {
  it('Visits the Homepage', () => {
    cy.visit('http://localhost:4200/')

    cy.contains('AlliRater').click()

    // Should be on a new URL which
    // includes '/homepage'
    cy.url().should('include', '/homepage')
  })
})