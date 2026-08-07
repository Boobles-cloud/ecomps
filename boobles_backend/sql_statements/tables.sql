-- So the program can check if the database and the app is on the newest version
CREATE TABLE Versions(
    VersionId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    VersionAppName varchar(50),
    VersionAppNum varchar(50),
    VersionDatabaseName varchar(50),
    VersionDatabaseNum varchar(50)
);

CREATE TABLE Languages(
    LangId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    LangName varchar(40),
    Country varchar(40)
);


-- All users for a tenant
CREATE TABLE Users(
    UserId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    UserName varchar(50) NOT NULL,
    UserPW varchar(100) NOT NULL,
    UserMail varchar(100) NOT NULL,
    UserTel varchar(50),
    UserHas2Fa bool,
    UserHasTenant bool,
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
    TenantActionId int unsigned NOT NULL,
    UserId int unsigned NOT NULL,

    FOREIGN KEY(TenantActionId) REFERENCES TenantActions(ActionId),
    FOREIGN KEY(UserId) REFERENCES Users(UserId)
);


-- ############################################################
--                      Tenant Configuration Stuff
-- ############################################################

-- The tenant all users are connected to
CREATE TABLE Tenant(
    TenantId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TenantName varchar(100) NOT NULL,
    TenantCreation DATETIME NOT NULL,
    TenantAdminUserId int unsigned NOT NULL,
    TenantPwId int unsigned NOT NULL,
    -- TODO: Add more tenant data
    FOREIGN KEY(TenantPwId) REFERENCES TenantPw(TenantPwId),
    FOREIGN KEY(TenantAdminUserId) REFERENCES Users(UserId)
);


-- Table to store all actions and permissions a tenant has
CREATE TABLE TenantActions(
    ActionId int unsigned NOT NULL,
    ActionName varchar(50) NOT NULL,
    ActionDescription varchar(50) NOT NULL,
    LanguageId int unsigned NOT NULL,

    FOREIGN KEY(LanguageId) REFERENCES Languages(LangId)
);

-- For storring the tenant pw
CREATE TABLE TenantPw(
    TenantPwId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TenantPwVal varchar(255)
);

-- Here are all tenants storred that are free for deletion
CREATE TABLE TenantDelitions(
    TenantDelitionId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    IssuedFrom varchar(250) NOT NULL,
    IssuedOn DATETIME,
    WhenToComplete DATETIME,
    Deleted bool,
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);


-- All tokens for applications like: Ebay, Amazon, etc.
CREATE TABLE TenantOAuthTokens(
    TokenId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TokenVal varchar(250) NOT NULL,
    LasChanget DATETIME,
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
    AppSyncActivated DATETIME NOT NULL,
    AppSync BOOL,
    AppSyncInterval int unsigned,
    TenantId int unsigned NOT NULL,

    FOREIGN KEY(TenantId) REFERENCES Tenant(TenantId)
);



-- ############################################################
--                         Tenant data stuff
--                       Like orders, Customers
-- ############################################################

-- All customers for a specific tenant
CREATE TABLE Customer(
    CustomerId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    CustomerName varchar(200) NOT NULL,
    CustomerPostalCode varchar(50),
    CustomerStreetAndHouseNr varchar(200),
    CustomerCity varchar(100),
    CustomerLastChanged DATETIME,
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
    OrderLastChanged DATETIME,

    FOREIGN KEY(OrderStatus) REFERENCES OrderStatus(StatusId)
);

-- To indicate which state an order has
CREATE TABLE OrderStatus(
    StatusId int unsigned NOT NULL AUTO_INCREMENT,
    StatusName varchar(200) NOT NULL,
    LanguageId int unsigned NOT NULL,

    FOREIGN KEY(LanguageId) REFERENCES Languages(LangId)
);


-- All products an order contains
CREATE TABLE OrderProducts(
    OPId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    ProductId int unsigned NOT NULL,
    Amount int unsigned NOT NULL,
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

-- For storring all product pictures
CREATE TABLE ProductPictures(
    PictureId int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    PicturePath varchar(250) NOT NULL,
    PicturePosition int unsigned,
    ProductId int unsigned NOT NULL,

    FOREIGN KEY(ProductId) REFERENCES Product(ProductId)
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



-- ############################################################
--                      Development Stuff
-- ############################################################


INSERT INTO Languages VALUES
    (1, "Deutsch", "Deutschland"),
    (2, "English GB", "Great Britin"),
    (3, "Svenska", "Sverige");


-- To set the default version
-- TODO: Change version here, before releasing a new version!
INSERT INTO Versions VALUES(DEFAULT, 'Alpha', '0.1', 'Alpha', '0.1');

-- We set the tenant default
-- We do this, so we can filter if the user has a tenant or not :)
INSERT INTO Tenant(TenantId, TenantName) VALUES(0, "USER_HAS_NO_TENANT");

-- Status stuff
INSERT INTO OrderStatus() VALUES
-- -------------------------------------------------------------
-- Deutsch (LanguageId: 1)
-- -------------------------------------------------------------
(1, 'Offen / Ausstehend', 1),
(2, 'Zahlung ausstehend', 1),
(3, 'Zahlung erhalten', 1),
(4, 'In Bearbeitung', 1),
(5, 'Versandbereit', 1),
(6, 'Versandt', 1),
(7, 'Zugestellt', 1),
(8, 'Abgeschlossen', 1),
(9, 'Storniert', 1),
(10, 'Rückabwicklung / Retoure', 1),

-- -------------------------------------------------------------
-- English GB (LanguageId: 2)
-- -------------------------------------------------------------
(1, 'Pending', 2),
(2, 'Payment Pending', 2),
(3, 'Payment Received', 2),
(4, 'Processing', 2),
(5, 'Ready for Dispatch', 2),
(6, 'Dispatched', 2),
(7, 'Delivered', 2),
(8, 'Completed', 2),
(9, 'Cancelled', 2),
(10, 'Returned', 2),

-- -------------------------------------------------------------
-- Svenska (LanguageId: 3)
-- -------------------------------------------------------------
(1, 'Väntande', 3),
(2,'Väntar på betalning', 3),
(3, 'Betalning mottagen', 3),
(4, 'Behandlas', 3),
(5, 'Redo för skickas', 3),
(6, 'Skickad', 3),
(7, 'Levererad', 3),
(8, 'Slutförd', 3),
(9, 'Avbruten', 3),
(10, 'Returnerad', 3);


-- All Actions a tenant can do
INSERT INTO TenantActions VALUES
-- -------------------------------------------------------------
-- Deutsch (LanguageId = 1)
-- -------------------------------------------------------------
(1, "AddUser", "Berechtigung, einen neuen Benutzer hinzuzufügen", 1),
(2, "AddOAuthApp", "Berechtigung, Drittanbieter-Apps zum Synchro hinzuzufügen", 1),
(3, "AddCustomer", "Berechtigung, einen neuen Kunden hinzuzufügen", 1),
(4, "AddOrder", "Berechtigung, eine neue Bestellung hinzuzufügen", 1),
(5, "AddOrderStatus", "Berechtigung, einen neuen Bestellstatus hinzuzufügen", 1),
(6, "AddProduct", "Berechtigung, ein neues Produkt hinzuzufügen", 1),
(7, "AddWarehouse", "Berechtigung, ein neues Lager hinzuzufügen", 1),
(8, "EditUserPermission", "Berechtigung, Benutzerrechte zu bearbeiten", 1),
(9, "EditTenantConfiguration", "Berechtigung, den Mandanten zu bearbeiten", 1),
(10, "EditOAuthApp", "Berechtigung, Drittanbieter-Apps zu bearbeiten", 1),
(11, "EditCustomer", "Berechtigung, einen Kunden zu bearbeiten", 1),
(12, "EditOrder", "Berechtigung, eine Bestellung zu bearbeiten", 1),
(13, "EditProduct", "Berechtigung, ein Produkt zu bearbeiten", 1),
(14, "EditWarehouse", "Berechtigung, ein Lager zu bearbeiten", 1),
(15, "DeleteUser", "Berechtigung, einen Benutzer zu löschen", 1),
(16, "DeleteOAuthApp", "Berechtigung, eine Drittanbieter-App zu löschen", 1),
(17, "DeleteCustomer", "Berechtigung, einen Kunden zu löschen", 1),
(18, "DeleteOrder", "Berechtigung, eine Bestellung zu löschen", 1),
(19, "DeleteProduct", "Berechtigung, ein Produkt zu löschen", 1),
(20, "DeleteWarehouse", "Berechtigung, ein Lager zu löschen", 1),

-- -------------------------------------------------------------
-- English GB (LanguageId = 2)
-- -------------------------------------------------------------
(1, "AddUser", "Permission to add a new user to this tenant", 2),
(2, "AddOAuthApp", "Permission to add new third party applications, to sync data", 2),
(3, "AddCustomer", "Permission to add a new customer", 2),
(4, "AddOrder", "Permission to add new order", 2),
(5, "AddOrderStatus", "Permission to add new order status", 2),
(6, "AddProduct", "Permission to add new product", 2),
(7, "AddWarehouse", "Permission to add a new warehouse", 2),
(8, "EditUserPermission", "Permission to edit a users permissions", 2),
(9, "EditTenantConfiguration", "Permission to edit the tenant", 2),
(10, "EditOAuthApp", "Permission to edit third party application", 2),
(11, "EditCustomer", "Permission to edit a customer", 2),
(12, "EditOrder", "Permission to edit an order", 2),
(13, "EditProduct", "Permission to edit a product", 2),
(14, "EditWarehouse", "Permission to edit a warehouse", 2),
(15, "DeleteUser", "Permission to delete a user", 2),
(16, "DeleteOAuthApp", "Permission to delete an third party application", 2),
(17, "DeleteCustomer", "Permission to delete a customer", 2),
(18, "DeleteOrder", "Permission to delete an order", 2),
(19, "DeleteProduct", "Permission to delete a product", 2),
(20, "DeleteWarehouse", "Permission to delete a warehouse", 2),

-- -------------------------------------------------------------
-- Svenska / Schwedisch (LanguageId = 3)
-- -------------------------------------------------------------
(1, "AddUser", "Behörighet att lägga till en ny användare", 3),
(2, "AddOAuthApp", "Behörighet att lägga till tredjepartsapplikationer", 3),
(3, "AddCustomer", "Behörighet att lägga till en ny kund", 3),
(4, "AddOrder", "Behörighet att lägga till en ny order", 3),
(5, "AddOrderStatus", "Behörighet att lägga till en ny orderstatus", 3),
(6, "AddProduct", "Behörighet att lägga till en ny produkt", 3),
(7, "AddWarehouse", "Behörighet att lägga till ett nytt lager", 3),
(8, "EditUserPermission", "Behörighet att redigera användarbehörigheter", 3),
(9, "EditTenantConfiguration", "Behörighet att redigera organisationen", 3),
(10, "EditOAuthApp", "Behörighet att redigera tredjepartsapplikation", 3),
(11, "EditCustomer", "Behörighet att redigera en kund", 3),
(12, "EditOrder", "Behörighet att redigera en order", 3),
(13, "EditProduct", "Behörighet att redigera en produkt", 3),
(14, "EditWarehouse", "Behörighet att redigera ett lager", 3),
(15, "DeleteUser", "Behörighet att ta bort en användare", 3),
(16, "DeleteOAuthApp", "Behörighet att ta bort en tredjepartsapplikation", 3),
(17, "DeleteCustomer", "Behörighet att ta bort en kund", 3),
(18, "DeleteOrder", "Behörighet att ta bort en order", 3),
(19, "DeleteProduct", "Behörighet att ta bort en produkt", 3),
(20, "DeleteWarehouse", "Behörighet att ta bort ett lager", 3);
