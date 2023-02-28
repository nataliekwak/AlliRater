describe('Homepage Test', () => {
  it('Visits the Homepage', () => {
    cy.visit('http://localhost:4200/')

    cy.contains('Home').click()

    // Should be on a new URL which
    // includes '/homepage'
    cy.url().should('include', '/homepage')
  })
})