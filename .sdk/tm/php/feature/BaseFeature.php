<?php
declare(strict_types=1);

// Newton SDK base feature

class NewtonBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(NewtonContext $ctx, array $options): void {}
    public function PostConstruct(NewtonContext $ctx): void {}
    public function PostConstructEntity(NewtonContext $ctx): void {}
    public function SetData(NewtonContext $ctx): void {}
    public function GetData(NewtonContext $ctx): void {}
    public function GetMatch(NewtonContext $ctx): void {}
    public function SetMatch(NewtonContext $ctx): void {}
    public function PrePoint(NewtonContext $ctx): void {}
    public function PreSpec(NewtonContext $ctx): void {}
    public function PreRequest(NewtonContext $ctx): void {}
    public function PreResponse(NewtonContext $ctx): void {}
    public function PreResult(NewtonContext $ctx): void {}
    public function PreDone(NewtonContext $ctx): void {}
    public function PreUnexpected(NewtonContext $ctx): void {}
}
