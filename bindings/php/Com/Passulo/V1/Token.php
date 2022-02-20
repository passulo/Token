<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/passulo/token.proto

namespace Com\Passulo\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>com.passulo.v1.Token</code>
 */
class Token extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string firstName = 2;</code>
     */
    protected $firstName = '';
    /**
     * Generated from protobuf field <code>string middleName = 3;</code>
     */
    protected $middleName = '';
    /**
     * Generated from protobuf field <code>string lastName = 4;</code>
     */
    protected $lastName = '';
    /**
     * Generated from protobuf field <code>string gender = 5;</code>
     */
    protected $gender = '';
    /**
     * Generated from protobuf field <code>string number = 6;</code>
     */
    protected $number = '';
    /**
     * Generated from protobuf field <code>string status = 7;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string company = 8;</code>
     */
    protected $company = '';
    /**
     * Generated from protobuf field <code>string email = 9;</code>
     */
    protected $email = '';
    /**
     * Generated from protobuf field <code>string telephone = 10;</code>
     */
    protected $telephone = '';
    /**
     * Generated from protobuf field <code>string association = 11;</code>
     */
    protected $association = '';
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp validUntil = 12;</code>
     */
    protected $validUntil = null;
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp memberSince = 13;</code>
     */
    protected $memberSince = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $firstName
     *     @type string $middleName
     *     @type string $lastName
     *     @type string $gender
     *     @type string $number
     *     @type string $status
     *     @type string $company
     *     @type string $email
     *     @type string $telephone
     *     @type string $association
     *     @type \Google\Protobuf\Timestamp $validUntil
     *     @type \Google\Protobuf\Timestamp $memberSince
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Com\Passulo\Token::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string firstName = 2;</code>
     * @return string
     */
    public function getFirstName()
    {
        return $this->firstName;
    }

    /**
     * Generated from protobuf field <code>string firstName = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setFirstName($var)
    {
        GPBUtil::checkString($var, True);
        $this->firstName = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string middleName = 3;</code>
     * @return string
     */
    public function getMiddleName()
    {
        return $this->middleName;
    }

    /**
     * Generated from protobuf field <code>string middleName = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setMiddleName($var)
    {
        GPBUtil::checkString($var, True);
        $this->middleName = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string lastName = 4;</code>
     * @return string
     */
    public function getLastName()
    {
        return $this->lastName;
    }

    /**
     * Generated from protobuf field <code>string lastName = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setLastName($var)
    {
        GPBUtil::checkString($var, True);
        $this->lastName = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string gender = 5;</code>
     * @return string
     */
    public function getGender()
    {
        return $this->gender;
    }

    /**
     * Generated from protobuf field <code>string gender = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setGender($var)
    {
        GPBUtil::checkString($var, True);
        $this->gender = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string number = 6;</code>
     * @return string
     */
    public function getNumber()
    {
        return $this->number;
    }

    /**
     * Generated from protobuf field <code>string number = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 7;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string company = 8;</code>
     * @return string
     */
    public function getCompany()
    {
        return $this->company;
    }

    /**
     * Generated from protobuf field <code>string company = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setCompany($var)
    {
        GPBUtil::checkString($var, True);
        $this->company = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string email = 9;</code>
     * @return string
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * Generated from protobuf field <code>string email = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->email = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string telephone = 10;</code>
     * @return string
     */
    public function getTelephone()
    {
        return $this->telephone;
    }

    /**
     * Generated from protobuf field <code>string telephone = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setTelephone($var)
    {
        GPBUtil::checkString($var, True);
        $this->telephone = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string association = 11;</code>
     * @return string
     */
    public function getAssociation()
    {
        return $this->association;
    }

    /**
     * Generated from protobuf field <code>string association = 11;</code>
     * @param string $var
     * @return $this
     */
    public function setAssociation($var)
    {
        GPBUtil::checkString($var, True);
        $this->association = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp validUntil = 12;</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getValidUntil()
    {
        return $this->validUntil;
    }

    public function hasValidUntil()
    {
        return isset($this->validUntil);
    }

    public function clearValidUntil()
    {
        unset($this->validUntil);
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp validUntil = 12;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setValidUntil($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->validUntil = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp memberSince = 13;</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getMemberSince()
    {
        return $this->memberSince;
    }

    public function hasMemberSince()
    {
        return isset($this->memberSince);
    }

    public function clearMemberSince()
    {
        unset($this->memberSince);
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp memberSince = 13;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setMemberSince($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->memberSince = $var;

        return $this;
    }

}

