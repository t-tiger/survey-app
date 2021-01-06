module.exports = {
  testPathIgnorePatterns: ['<rootDir>/node_modules/'],
  modulePaths: ['<rootDir>'],
  moduleNameMapper: {
    '^components/(.*)$': '<rootDir>/components/$1',
    '^modules/(.*)$': '<rootDir>/modules/$1',
    '^const/(.*)$': '<rootDir>/const/$1',
    '^util/(.*)$': '<rootDir>/util/$1',
  },
}
