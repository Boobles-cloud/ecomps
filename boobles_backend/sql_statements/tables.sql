-- So the program can check if the database and the app is on the newest version
CREATE TABLE Versions(
    VersionId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    VersionAppName varchar(50),
    VersionAppNum varchar(50),
    VersionDatabaseName varchar(50),
    VersionDatabaseNum varchar(50)
);

-- All users for a tenant
CREATE TABLE Users(
    UserId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    UserName varchar(50) NOT NULL,
    UserPW varchar(100) NOT NULL,
    UserMail varchar(100) NOT NULL,
    UserTel varchar(50),
    UserHas2Fa bool,
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);

-- To store all accesstokens for the users
CREATE TABLE UserAccesstokens(
    UserAccessId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TokenVal varchar(250) NOT NULL,
    TokenExpire DATETIME NOT NULL,
    UserId int unsigned,

    FOREIGN KEY(UserId) REFERENCES Users(UserId)
);

-- For all permissions a user can have
-- Like Admin, etc.
CREATE TABLE Permissions(
    PermissionId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    PermissionName varchar(100),
    PermissionDescription varchar(150),
    UserId int unsigned NOT NULL,

    FOREIGN KEY(UserId) REFERENCES Users(UserId)
);

-- The tenant all users are connected to
CREATE TABLE Tenant(
    TenantId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TenantName varchar(100) NOT NULL,
    TenantPw varchar(250) NOT NULL -- The tenant PW (or key) where all data gets encrypted by it -> generated automaticly on setup?
    -- TODO: Add more tenant data
);

-- All tokens for applications like: Ebay, Amazon, etc.
CREATE TABLE TenantOAuthTokens(
    TokenId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TokenVal varchar(250) NOT NULL,
    AppId int unsigned NOT NULL,
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(AppId) REFERENCES OAuthApplications(AppId),
    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);

-- All capable oauth applications
-- Also all metadata for syncs -> We sync all data
CREATE TABLE OAuthApplications(
    AppId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    AppName varchar(150) NOT NULL,
    AppStruct varchar(150) NOT NULL,
    AppSync BOOL,
    AppSyncInterval int unsigned,
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);

-- All customers for a specific tenant
CREATE TABLE Customer(
    CustomerId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    CustomerName varchar(200) NOT NULL,
    CustomerPostalCode varchar(50),
    CustomerStreetAndHouseNr varchar(200),
    CustomerCity varchar(100),
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);

-- All orders for a specific tenant
CREATE TABLE Order(
    OrderId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    OrderName varchar(250),
    OrderDate DATETIME NOT NULL,
    OrderStatus int unsigned NOT NULL,
    OrderPostalCode varchar(50), -- if the address isnt the same
    OrderStreetAndHouseNr varchar(200),
    OrderCity varchar(200),
);

-- To indicate which state an order has
CREATE TABLE OrderStatus(
    StatusId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    StatusName varchar(200) NOT NULL
);


-- All products an order contains
CREATE TABLE OrderProducts(
    OPId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ProductId int unsigned NOT NULL,
    OrderId int unsigned NOT NULL,

    FOREIGN KEY(ProductId) REFERENCES Product(ProductId),
    FOREIGN KEY(OrderId) REFERENCES Order(OrderId)
);

-- All products a tenant has
CREATE TABLE Product(
    ProductId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ProductName varchar(250) NOT NULL,
    ProductPrice varchar(250) NOT NULL,
    ProductDescription varchar(250) NOT NULL,
    ProductPicturePath varchar(100),
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);

-- If a Product is storred in multiple warehouses
CREATE TABLE ProductWarehouses(
    ProdWareId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Amount int unsigned NOT NULL,
    ProductId int unsigned NOT NULL,
    WarehouseId int unsigned NOT NULL,

    FOREIGN KEY(ProductId) REFERENCES Product(ProductId),
    FOREIGN KEY(WarehouseId) REFERENCES Warehouse(WarehouseId)
);

-- All warehouses a tenant has
CREATE TABLE Warehouse(
    WarehousId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    WarehousName varchar(150) NOT NULL,
    WarehousePostalCode varchar(50),
    WarehouseHouseNumAndStreet varchar(250),
    WarehouseCity varchar(250),
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);

-- To set the default version
-- TODO: Change version here, before releasing a new version!
INSERT INTO Versions VALUES(DEFAULT, 'Alpha', '0.1', 'Alpha', '0.1');
