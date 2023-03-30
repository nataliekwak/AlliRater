describe('Webpage Test', () => {
  it('Test\'s all functional domains', () => {
    cy.visit('http://localhost:4200/')

    cy.contains('Dashboard').click()

    cy.wait(500)

    cy.url().should('include', '/#')
    
    cy.wait(500)

    cy.visit('http://localhost:4200/login')

    cy.wait(500)

    cy.url().should('include', '/login')

    cy.wait(500)

    cy.visit('http://localhost:4200/register')

    cy.wait(500)

    cy.url().should('include', '/register')
  })
})