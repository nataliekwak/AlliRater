describe('Sprint 2 First Test', () => {
  it('Visits the Homepage', () => {
    cy.visit('http://localhost:4200/')

    cy.contains('Homepage').click()

    // Should be on a new URL which
    // includes '/commands/actions'
    cy.url().should('include', '/homepage')
  })
})